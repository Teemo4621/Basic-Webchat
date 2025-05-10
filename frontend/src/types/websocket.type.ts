export type WebSocketMessage = {
    method: "system" | "message";
    data: {
      id?: number;
      message_id: string;
      room_id: number;
      content: string;
      created_at: string;
      user_id: number;
      username: string;
      online?: boolean;
    };
  };