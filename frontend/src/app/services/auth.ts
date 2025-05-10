import { LoginForm, RegisterForm } from "@/types/auth.type";
import AxiosInstance from "../../libs/axiosInstance";
import { FetchLoginResponse, FetchRegisterResponse } from "@/types/auth.type";
import { FetchMeResponse } from "@/types/user.type";
import axios from "axios";

const AuthServices = {
    login: async ({ username, password }: LoginForm): Promise<FetchLoginResponse> => {
        try {
            const response = await AxiosInstance.post("/auth/login", {
                username,
                password,
            });
            return response.data;
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            }
            throw new Error("Unexpected error");
        }
    },

    register: async ({ username, password }: RegisterForm): Promise<FetchRegisterResponse> => {
        try {
            const response = await AxiosInstance.post("/auth/register", {
                username,
                password,
            });
            return response.data;
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            }
            throw new Error("Unexpected error");
        }
    },

    GetMe: async (): Promise<FetchMeResponse> => {
        try {
            const response = await AxiosInstance.get<FetchMeResponse>("/auth/@me");
            return response.data;
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            }
            throw new Error("Unexpected error");
        }
    },
}

export default AuthServices