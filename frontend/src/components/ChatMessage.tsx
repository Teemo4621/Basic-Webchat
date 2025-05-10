import { formatDistanceToNow } from "date-fns";
import { cn } from "@/libs/utils";
import { Message } from "@/types/messages";
import { Member } from "@/types/rooms.type";

interface ChatMessageProps {
  message: Message;
  isOwnMessage?: boolean;
  user?: Member;
}

const ChatMessage = ({ message, isOwnMessage, user }: ChatMessageProps) => {
  const isSystem = message.sender === "system";
  const timeAgo = formatDistanceToNow(new Date(message.created_at), {
    addSuffix: true,
  });

  if (isSystem) {
    return (
      <div className="flex justify-center my-4 animate-fadeIn">
        <div className="bg-white/5 rounded-md px-3 py-2 text-sm text-white/60 max-w-[80%]">
          {message.content}
        </div>
      </div>
    );
  }

  return (
    <div
      className={cn(
        "flex mb-4 animate-fadeIn",
        isOwnMessage ? "justify-end" : "justify-start"
      )}
    >
      {!isOwnMessage && (
        <>
          <div className="w-8 h-8 rounded-full bg-white/10 flex items-center justify-center mr-2">
            {user?.username.charAt(0).toUpperCase()}
          </div>
        </>
      )}
      <div className="max-w-[80%]">
        {!isOwnMessage && (
          <p className="text-xs text-white/60 mb-1 font-medium">
            {user?.username}
          </p>
        )}
        <div
          className={cn(
            "w-full rounded-2xl px-4 py-2",
            isOwnMessage
              ? "bg-white text-black rounded-br-none"
              : "border border-white/10 glass rounded-tl-none"
          )}
        >
          {!isOwnMessage && (
            <div className="text-xs text-white/60 mb-1 font-medium">
              {message.sender}
            </div>
          )}
          <div className="break-words">{message.content}</div>
          <div
            className={cn(
              "text-[10px] mt-1 text-right",
              isOwnMessage ? "text-black/50" : "text-white/50"
            )}
          >
            {timeAgo}
          </div>
        </div>
      </div>
    </div>
  );
};

export default ChatMessage;
