export interface Message {
    id: number,
    message_id: string,
    room_id: number,
    content: string,
    created_at: string,
    user_id: number
    sender?: string
}

export interface SendMessageForm {
    content: string;
}

export interface DeleteMessageForm {
    message_code: string;
}

export interface DeleteMessageResponse {
    message: string,
    status: string
}

export interface FetchSendMessageResponse {
    data: Message,
    message: string,
    status: string
}

export interface FetchGetMessagesResponse {
    data: {
        messages: Message[];
        pagination: {
            limit: number,
            page: number,
            page_total: number,
        }
    }
}