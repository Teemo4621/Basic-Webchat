import axios from "axios";
import AxiosInstance from "../../libs/axiosInstance";
import { DeleteMessageForm, DeleteMessageResponse, FetchGetMessagesResponse, FetchSendMessageResponse, SendMessageForm } from "@/types/messages";

const MessagesServices = ({ room_code }: { room_code: string }) => {
    const Send = async ({ content }: SendMessageForm): Promise<FetchSendMessageResponse> => {
        try {
            const response = await AxiosInstance.post<FetchSendMessageResponse>(`/rooms/${room_code}/messages`, {
                content,
            });
            return response.data;
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            }
            throw new Error("Unexpected error");
        }
    };

    const GetList = async ({ limit, page }: { limit: number, page: number }): Promise<FetchGetMessagesResponse> => {
        try {
            const response = await AxiosInstance.get<FetchGetMessagesResponse>(`/rooms/${room_code}/messages?limit=${limit}&page=${page}`);
            return response.data;
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            }
            throw new Error("Unexpected error");
        }
    };

    const DeleteMessage = async ({ message_code }: DeleteMessageForm): Promise<DeleteMessageResponse> => {
        try {
            const response = await AxiosInstance.post<DeleteMessageResponse>(`/rooms/${room_code}/messages/${message_code}/delete`);
            return response.data;
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            }
            throw new Error("Unexpected error");
        }
    };

    return {
        Send,
        GetList,
        DeleteMessage,
    };
}

export default MessagesServices;