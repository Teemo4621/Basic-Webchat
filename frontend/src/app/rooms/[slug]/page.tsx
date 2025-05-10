"use client";

import { useEffect, useRef, useState } from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { ArrowLeft, Clipboard, SendIcon, Users } from "lucide-react";
import ChatMessage from "@/components/ChatMessage";
import UserList from "@/components/UserList";
import { useUserContext } from "@/contexts/userProvider";
import useRoomsHook from "@/hooks/roomsHook";
import useMessagesHook from "@/hooks/messagesHook";
import React from "react";
import { toast } from "sonner";
import debounce from "lodash.debounce";
import { motion, AnimatePresence } from "framer-motion";
import useWebSocket from "react-use-websocket";
import Cookies from "js-cookie";
import { WebSocketMessage } from "@/types/websocket.type";
import { Message } from "@/types/messages";

export default function ChatRoom({
  params,
}: {
  params: Promise<{ slug: string }>;
}) {
  const router = useRouter();
  const { slug } = React.use(params);
  const room_code = slug;

  const { user } = useUserContext();
  const [inputMessage, setInputMessage] = useState("");
  const [showUserList, setShowUserList] = useState(false);

  const messageEndRef = useRef<HTMLDivElement>(null);
  const messageContainerRef = useRef<HTMLDivElement>(null);

  const {
    room,
    getRoom,
    members,
    setMembers,
    getMembers,
    loading: roomsLoading,
    error: roomsError,
  } = useRoomsHook();
  const {
    messages,
    getMessages,
    sendMessage,
    setMessages,
    page,
    hasMore,
    loading: messagesLoading,
    error: messagesError,
  } = useMessagesHook();

  useEffect(() => {
    const limit = window.innerWidth >= 1024 ? 12 : 10;
    getRoom({ room_code: room_code });
    getMembers({ room_code: room_code });
    getMessages({ room_code: room_code, limit: limit, page: 1 });
  }, [room_code]);

  const token = Cookies.get("auth.token");
  const { sendJsonMessage, lastJsonMessage } = useWebSocket(
    `${process.env.NEXT_PUBLIC_API_URL?.replace(
      "http",
      "ws"
    )}/rooms/${room_code}/ws?token=${token}`,
    {
      onOpen: () => {
        console.log("WebSocket connected");
      },
      onError: (error) => {
        console.error("WebSocket error:", error);
        toast("WebSocket connection failed.");
      },
      onClose: () => {
        console.log("WebSocket disconnected");
      },
      shouldReconnect: () => true,
    }
  );

  useEffect(() => {
    if (!lastJsonMessage) return;
    const data = lastJsonMessage as WebSocketMessage;
    if (!data.method || !data.data) return;
    console.log(data);
    if (data.method === "system") {
      setMembers((prev) => {
        return prev.map((m) =>
          m.username == data.data.username
            ? { ...m, online: data.data.online }
            : m
        );
      });
    } else if (
      data.method === "message" &&
      data.data.content &&
      data.data.message_id &&
      data.data.room_id &&
      data.data.user_id &&
      data.data.username
    ) {
      const newMessage = data.data;
      setMessages((prev) => {
        const exists = prev.some(
          (msg) => msg.message_id === newMessage.message_id
        );
        if (!exists) {
          const messageToAdd: Message = {
            id: newMessage.id || 0,
            message_id: newMessage.message_id,
            room_id: newMessage.room_id,
            content: newMessage.content,
            created_at: newMessage.created_at,
            user_id: newMessage.user_id,
          };
          return [messageToAdd, ...prev];
        }
        return prev;
      });
    }
  }, [lastJsonMessage]);

  useEffect(() => {
    const container = messageContainerRef.current;

    if (!container) return;

    if (page === 1) {
      messageEndRef.current?.scrollIntoView({ behavior: "smooth" });
    } else {
      container.scrollTop += 20;
    }
  }, [messages, page]);

  useEffect(() => {
    const handleScroll = debounce(() => {
      const limit = window.innerWidth >= 1024 ? 12 : 10;
      if (!messageContainerRef.current || messagesLoading || !hasMore) return;

      const container = messageContainerRef.current;
      const { scrollTop, scrollHeight } = container;

      if (scrollTop === 0) {
        const previousScrollHeight = scrollHeight;

        getMessages({
          room_code: room_code,
          limit: limit,
          page: page + 1,
        }).then(() => {
          if (messageContainerRef.current) {
            const newScrollHeight = messageContainerRef.current.scrollHeight;
            messageContainerRef.current.scrollTop =
              newScrollHeight - previousScrollHeight;
          }
        });
      }
    }, 200);

    const container = messageContainerRef.current;
    container?.addEventListener("scroll", handleScroll);

    return () => {
      container?.removeEventListener("scroll", handleScroll);
      handleScroll.cancel();
    };
  }, [page, room_code, messagesLoading, hasMore, getMessages]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (!inputMessage.trim() || !user || !user.id) {
      toast("Please log in to send messages.");
      return;
    }
    sendMessage({ room_code: room_code, content: inputMessage, wsSendMsg: sendJsonMessage });
    setInputMessage("");
    messageEndRef.current?.scrollIntoView({ behavior: "smooth" });
  };

  const handleBack = () => {
    router.push("/rooms");
  };

  const handleCopyRoomCode = () => {
    navigator.clipboard.writeText(room_code);
    toast("Room code copied to clipboard");
  };

  if (roomsLoading || !room || !members) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="backdrop-blur-xl border border-white/10 p-6 rounded-xl">
          <p className="text-white/80">Loading your chat room...</p>
        </div>
      </div>
    );
  }

  if (roomsError || messagesError) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="backdrop-blur-xl border border-white/10 p-6 rounded-xl">
          <p className="text-red-500">{roomsError || messagesError}</p>
        </div>
      </div>
    );
  }

  return (
    <div className="flex flex-col h-screen">
      <header className="glass border-b border-white/10 px-4 py-3 flex items-center justify-between">
        <div className="flex items-center">
          <Button
            variant="ghost"
            size="icon"
            onClick={handleBack}
            className="mr-2 hover:bg-white/5"
          >
            <ArrowLeft className="h-5 w-5" />
          </Button>
          <div>
            <h1 className="font-bold">Room: {room?.name}</h1>
            <p className="text-xs text-white/60">
              Joined as {user?.username || "Guest"}
            </p>
          </div>
        </div>
        <div>
          <Button
            variant={showUserList ? "secondary" : "ghost"}
            size="icon"
            onClick={handleCopyRoomCode}
            className="hover:bg-white/5"
          >
            <Clipboard className="h-5 w-5" />
          </Button>
          <Button
            variant={showUserList ? "secondary" : "ghost"}
            size="icon"
            onClick={() => setShowUserList(!showUserList)}
            className="hover:bg-white/5"
          >
            <Users className="h-5 w-5" />
          </Button>
        </div>
      </header>

      <div className="flex flex-1 overflow-hidden">
        <div
          ref={messageContainerRef}
          className="flex-1 overflow-y-scroll px-4 pt-4 space-y-4"
          role="log"
          aria-live="polite"
        >
          {messagesLoading && !messages.length && (
            <div className="text-center text-white/60">Loading messages...</div>
          )}
          {!messagesLoading && !messages.length && (
            <div className="text-center text-white/60">No messages yet.</div>
          )}
          <AnimatePresence>
            {messages.toReversed().map((message) => (
              <motion.div
                key={message.message_id}
                initial={{ opacity: 0, y: -10 }}
                animate={{ opacity: 1, y: 0 }}
                exit={{ opacity: 0, y: -10 }}
                transition={{ duration: 0.5 }}
              >
                <ChatMessage
                  message={message}
                  isOwnMessage={message.user_id === user?.id}
                  user={members.find(
                    (member) => Number(member.user_id) === message.user_id
                  )}
                />
              </motion.div>
            ))}
            <div ref={messageEndRef} />
          </AnimatePresence>
        </div>

        {showUserList && (
          <div className="w-64 glass border-l border-white/10 overflow-y-auto animate-slideIn">
            <UserList users={members} room={room} />
          </div>
        )}
      </div>

      <form
        onSubmit={handleSubmit}
        className="p-4 glass border-t border-white/10"
      >
        <div className="flex items-center space-x-2">
          <Input
            type="text"
            placeholder="Type a message..."
            className="flex-1 bg-white/5 border-white/10 focus:border-white/30"
            value={inputMessage}
            onChange={(e) => setInputMessage(e.target.value)}
            aria-label="Message input"
          />
          <Button
            type="submit"
            className="bg-white text-black hover:bg-white/90"
            disabled={!inputMessage.trim() || !user}
          >
            <SendIcon className="h-4 w-4" />
          </Button>
        </div>
      </form>
    </div>
  );
}
