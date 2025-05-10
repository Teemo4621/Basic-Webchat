import MessagesServices from "@/app/services/messages";
import { useState } from "react";
import { isAxiosError } from "axios";
import { toast } from "sonner";
import { Message } from "@/types/messages";
import { useUserContext } from "@/contexts/userProvider";
import { SendJsonMessage } from "react-use-websocket/dist/lib/types";
import { WebSocketMessage } from "@/types/websocket.type";

const useMessagesHook = () => {
  const [messages, setMessages] = useState<Message[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<boolean>(false);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(true);

  const { user } = useUserContext();

  const resetState = () => {
    setLoading(false);
    setError(null);
    setSuccess(false);
  };

  const getMessages = async (data: {
    room_code: string;
    limit: number;
    page: number;
  }) => {
    if (data.page > 1 && !hasMore) return;
    resetState();
    setLoading(true);
    try {
      const response = await MessagesServices({
        room_code: data.room_code,
      }).GetList({
        limit: data.limit,
        page: data.page,
      });
      const newMessages = response.data.messages;
      setMessages((prev) => {
        const allMessages =
          data.page === 1 ? newMessages : [...prev, ...newMessages];
        const uniqueMessages = allMessages.filter(
          (msg, index, self) =>
            index === self.findIndex((m) => m.message_id === msg.message_id)
        );
        return uniqueMessages;
      });
      setSuccess(true);
      setPage(data.page);
      setHasMore(response.data.pagination.page_total > data.page);
    } catch (error) {
      setSuccess(false);
      if (isAxiosError(error)) {
        setError(error.response?.data?.message || "Get messages failed.");
        toast(error.response?.data?.message || "Get messages failed.");
      } else {
        setError("Unexpected error.");
        toast("Unexpected error.");
      }
    } finally {
      setLoading(false);
    }
  };

  const sendMessage = async (data: {
    room_code: string;
    content: string;
    wsSendMsg: SendJsonMessage;
  }) => {
    resetState();
    setLoading(true);
    try {
      const { data: message } = await MessagesServices({
        room_code: data.room_code,
      }).Send({
        content: data.content,
      });
      setMessages((prev) => [message, ...prev]);
      setSuccess(true);
      const wsMessage: WebSocketMessage = {
        method: "message",
        data: {
          message_id: message.message_id,
          room_id: message.room_id,
          content: message.content,
          created_at: message.created_at,
          user_id: message.user_id,
          username: user?.username || "Guest",
        },
      };
      data.wsSendMsg(wsMessage);
    } catch (error) {
      setSuccess(false);
      if (isAxiosError(error)) {
        setError(error.response?.data?.message || "Send message failed.");
        toast(error.response?.data?.message || "Send message failed.");
      } else {
        setError("Unexpected error.");
        toast("Unexpected error.");
      }
    } finally {
      setLoading(false);
    }
  };

  return {
    messages,
    loading,
    error,
    success,
    page,
    hasMore,
    getMessages,
    setMessages,
    sendMessage,
    setPage,
  };
};

export default useMessagesHook;
