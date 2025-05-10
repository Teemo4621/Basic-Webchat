import axios, { AxiosError } from "axios";
import Cookies from "js-cookie";
const AxiosInstance = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL,
    headers: {
        "Content-Type": "application/json",
    },
    timeout: 10 * 1000,
});

AxiosInstance.interceptors.request.use(
    (config) => {
        const token = Cookies.get("auth.token") ?? null;
        if (token) {
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
);

AxiosInstance.interceptors.response.use(
    (response) => {
        // when login success
        if (response.config.url === "/auth/login") {
          if (!response.data.data.access_token || !response.data.data.refresh_token)
            return response;
          Cookies.set("auth.token", response.data.data.access_token);
          Cookies.set("auth.reftoken", response.data.data.refresh_token);
        }
        return response;
    },
    async (error: AxiosError) => {
        const originalRequest = error.config;

        //when token expired
        if (error.response?.status === 401 && originalRequest && originalRequest.url !== "/auth/refresh-token") {
            try {
                const refreshToken = Cookies.get("auth.reftoken") ?? "";

                //when refresh token not found
                if (!refreshToken) {
                    Cookies.remove("auth.token");
                    Cookies.remove("auth.reftoken");
                    return Promise.reject(error);
                }

                const response = await AxiosInstance.post("/auth/refresh-token", {
                    refresh_token: refreshToken,
                });

                //when refresh token success
                if (response.data.data.access_token && response.data.data.refresh_token) {
                    Cookies.set("auth.token", response.data.data.access_token);
                    Cookies.set("auth.reftoken", response.data.data.refresh_token);
                    originalRequest.headers.Authorization = `Bearer ${response.data.data.access_token}`;
                    return AxiosInstance(originalRequest);
                }
            } catch (error) {
                //when refresh token failed
                Cookies.remove("auth.token");
                Cookies.remove("auth.reftoken");
                return Promise.reject(error);
            }
        }
        return Promise.reject(error);
    }
);

export default AxiosInstance;