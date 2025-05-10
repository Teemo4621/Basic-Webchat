import { CreateRoomResponse, DeleteRoomResponse, FetchJoinResponse, FetchRoomListResponse, FetchRoomResponse, GetMembersResponse, LeaveRoomResponse, RoomForm } from "@/types/rooms.type";
import AxiosInstance from "../../libs/axiosInstance";
import axios from "axios";

const RoomServices = {
    Create: async ({ name, description }: RoomForm): Promise<CreateRoomResponse> => {
        try {
            const response = await AxiosInstance.post<CreateRoomResponse>("/rooms", {
                name,
                description,
            });
            return response.data;
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            }
            throw new Error("Unexpected error");
        }
    },

    GetList: async ({ limit, page }: { limit: number, page: number }): Promise<FetchRoomListResponse> => {
        try {
            const response = await AxiosInstance.get<FetchRoomListResponse>(`/rooms?limit=${limit}&page=${page}`);
            return response.data;
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            }
            throw new Error("Unexpected error");
        }
    },

    Get: async ({ room_code }: { room_code: string }): Promise<FetchRoomResponse> => {
        try {
            const response = await AxiosInstance.get<FetchRoomResponse>(`/rooms/${room_code}`);
            return response.data;
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            }
            throw new Error("Unexpected error");
        }
    },

    Join: async ({ room_code }: { room_code: string }): Promise<FetchJoinResponse> => {
        try {
            const response = await AxiosInstance.post<FetchJoinResponse>(`/rooms/${room_code}/join`);
            return response.data;
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            }
            throw new Error("Unexpected error");
        }
    },

    Leave: async ({ room_code }: { room_code: string }): Promise<LeaveRoomResponse> => {
        try {
            const response = await AxiosInstance.post<LeaveRoomResponse>(`/rooms/${room_code}/leave`);
            return response.data;
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            }
            throw new Error("Unexpected error");
        }
    },

    Delete: async ({ room_code }: { room_code: string }): Promise<DeleteRoomResponse> => {
        try {
            const response = await AxiosInstance.post<DeleteRoomResponse>(`/rooms/${room_code}/delete`);
            return response.data;
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            }
            throw new Error("Unexpected error");
        }
    },

    GetMembers: async ({ room_code }: { room_code: string }): Promise<GetMembersResponse> => {
        try {
            const response = await AxiosInstance.get<GetMembersResponse>(`/rooms/${room_code}/members`);
            return response.data;
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            }
            throw new Error("Unexpected error");
        }
    },
};

export default RoomServices;
