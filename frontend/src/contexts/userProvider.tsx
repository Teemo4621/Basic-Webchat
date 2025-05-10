"use client";

import AuthServices from "@/app/services/auth";
import { User } from "@/types/user.type";
import { createContext, useContext, useEffect, useState } from "react";
import Cookies from "js-cookie";

interface UserContextProps {
  user: User | null;
}

export const UserContext = createContext<UserContextProps | null>(null);

export const UserProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [user, setUser] = useState<User | null>(null);
  
  const accessToken = Cookies.get("auth.token") ?? null;
  useEffect(() => {
    const fetchUser = async () => {
      try {
        const { data } = await AuthServices.GetMe();
        setUser(data);
      } catch (error) {
        console.error("Error fetching user:", error);
      }
    };
    
    if (accessToken) {
      fetchUser();
    }
  }, [accessToken]);

  return (
    <UserContext.Provider value={{ user }}>
      {children}
    </UserContext.Provider>
  );
};

export const useUserContext = () => {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error("useUserContext must be used within a UserProvider");
  }
  return context;
};
