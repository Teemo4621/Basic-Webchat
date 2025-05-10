import RoomServices from "@/app/services/rooms";
import { Member, Room } from "@/types/rooms.type";
import { isAxiosError } from "axios";
import { useState } from "react";
import { toast } from "sonner";

const useRoomsHook = () => {
    const [room, setRoom] = useState<Room | null>(null);
    const [rooms, setRooms] = useState<Room[]>([]);
    const [members, setMembers] = useState<Member[]>([]);
    const [success, setSuccess] = useState<boolean>(false);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const getRoom = async (data: { room_code: string }) => {
        resetState();
        try {
            const { data: room } = await RoomServices.Get(data);
            setRoom(room);
            setSuccess(true);
        } catch (error) {
            if (isAxiosError(error)) {
                setError(error.response?.data?.message || "Get room failed.");
                toast(error.response?.data?.message || "Get room failed.");
            } else {
                setError("Unexpected error.");
                toast("Unexpected error.");
            }
            setSuccess(false);
        } finally {
            setLoading(false);
        }
    };

    const getRooms = async () => {
        resetState();
        try {
            const { data } = await RoomServices.GetList({ limit: 10, page: 1 });
            setRooms(data.rooms);
            setSuccess(true);
        } catch (error) {
            if (isAxiosError(error)) {
                setError(error.response?.data?.message || "Get rooms failed.");
                toast(error.response?.data?.message || "Get rooms failed.");
            } else {
                setError("Unexpected error.");
                toast("Unexpected error.");
            }
            setSuccess(false);
        } finally {
            setLoading(false);
        }
    };

    const createRoom = async (data: { name: string, description: string }) => {
        resetState();
        try {
            const { data: room } = await RoomServices.Create(data);
            toast("Create room successful.");
            setRooms((prevRooms) => [...prevRooms, room]);
            setSuccess(true);
        } catch (error) {
            if (isAxiosError(error)) {
                setError(error.response?.data?.message || "Create room failed.");
                toast(error.response?.data?.message || "Create room failed.");
            } else {
                setError("Unexpected error.");
                toast("Unexpected error.");
            }
            setSuccess(false);
        } finally {
            setLoading(false);
        }
    };

    const deleteRoom = async (data: { room_code: string }) => {
        resetState();
        try {
            await RoomServices.Delete(data);
            toast("Delete room successful.");
            setRooms((prevRooms) => prevRooms.filter((room) => room.room_code !== data.room_code));
            setSuccess(true);
        } catch (error) {
            setSuccess(false);
            if (isAxiosError(error)) {
                setError(error.response?.data?.message || "Delete room failed.");
                toast(error.response?.data?.message || "Delete room failed.");
            } else {
                setError("Unexpected error.");
                toast("Unexpected error.");
            }
        } finally {
            setLoading(false);
        }
    };

    const joinRoom = async (data: { room_code: string }) => {
        resetState();
        try {
            await RoomServices.Join(data);
            toast("Join room successful.");
            setSuccess(true);
        } catch (error) {
            setSuccess(false);
            if (isAxiosError(error)) {
                setError(error.response?.data?.message || "Join room failed.");
                toast(error.response?.data?.message || "Join room failed.");
            } else {
                setError("Unexpected error.");
                toast("Unexpected error.");
            }
        } finally {
            setLoading(false);
        }
    };

    const leaveRoom = async (data: { room_code: string }) => {
        resetState();
        try {
            await RoomServices.Leave(data);
            setRooms((prevRooms) => prevRooms.filter((room) => room.room_code !== data.room_code));
            toast("Leave room successful.");
            setSuccess(true);
        } catch (error) {
            setSuccess(false);
            if (isAxiosError(error)) {
                setError(error.response?.data?.message || "Leave room failed.");
                toast(error.response?.data?.message || "Leave room failed.");
            } else {
                setError("Unexpected error.");
                toast("Unexpected error.");
            }
        } finally {
            setLoading(false);
        }
    };

    const getMembers = async (data: { room_code: string }) => {
        resetState();
        try {
            const { data: members } = await RoomServices.GetMembers(data);
            setMembers(members);
            setSuccess(true);
        } catch (error) {
            setSuccess(false);
            if (isAxiosError(error)) {
                setError(error.response?.data?.message || "Get members failed.");
                toast(error.response?.data?.message || "Get members failed.");
            } else {
                setError("Unexpected error.");
                toast("Unexpected error.");
            }
        } finally {
            setLoading(false);
        }
    };

    const resetState = () => {
        setLoading(false);
        setError(null);
        setSuccess(false);
    };

    return {
        rooms,
        members,
        room,
        loading,
        error,
        success,
        setMembers,
        getRoom,
        getRooms,
        createRoom,
        deleteRoom,
        joinRoom,
        leaveRoom,
        getMembers
    };
}

export default useRoomsHook;