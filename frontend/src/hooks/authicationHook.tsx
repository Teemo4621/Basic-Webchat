import AuthServices from "@/app/services/auth";
import { isAxiosError } from "axios";
import { useState } from "react";
import { toast } from "sonner";

const useAuthentication = () => {
    const [success, setSuccess] = useState<boolean>(false);
    const [loading, setLoading] = useState(false);  
    const [error, setError] = useState<string | null>(null);

    const login = async (data: { username: string, password: string }) => {
        setLoading(true);
        setError(null);
        setSuccess(false);
        try {
            await AuthServices.login(data);
            toast("Login successful.");
            setSuccess(true);
        } catch (error) {
            if (isAxiosError(error)) {
                setError(error.response?.data?.message || "Login failed.");
                toast(error.response?.data?.message || "Login failed.");
            } else {
                setError("Unexpected error.");
                toast("Unexpected error.");
            }
            setSuccess(false);
        } finally {
            setLoading(false);
        }
    };

    const register = async (data: { username: string, password: string }) => {
        setLoading(true);
        setError(null);
        setSuccess(false);
        try {
            await AuthServices.register(data);
            toast("Register successful.");
            setSuccess(true);
        } catch (error) {
            if (isAxiosError(error)) {
                setError(error.response?.data?.message || "Register failed.");
                toast(error.response?.data?.message || "Register failed.");
            } else {
                console.log(error);
                setError("Unexpected error.");
                toast("Unexpected error.");
            }
            setSuccess(false);
        } finally {
            setLoading(false);
        }
    };

    return {
        loading,
        error,
        success,
        login,
        register
    };
}

export default useAuthentication;
