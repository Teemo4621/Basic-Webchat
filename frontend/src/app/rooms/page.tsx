"use client"

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogFooter } from "@/components/ui/dialog";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@/components/ui/dropdown-menu";
import { Clock, MessageSquare, Plus, Trash2, ArrowRight, User, Settings, LogOut, DoorOpen } from "lucide-react";
import { ScrollArea } from "@/components/ui/scroll-area";
import { formatDate } from "@/libs/utils";
import Image from "next/image";
import useRoomsHook from "@/hooks/roomsHook";
import { useUserContext } from "@/contexts/userProvider";
import { toast } from "sonner";

const SelectChats = () => {
  const [roomCode, setRoomCode] = useState("");
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [roomName, setRoomName] = useState("");
  const [roomDescription, setRoomDescription] = useState("");

  const router = useRouter();
  const { rooms, getRooms, createRoom, loading, joinRoom, success, leaveRoom, deleteRoom } = useRoomsHook();
  const { user } = useUserContext();

  useEffect(() => {
    getRooms();
  }, []);

  const handleJoinRoom = (roomId: string) => {
    router.push(`/rooms/${roomId}`);
  };

  const handleCreateRoom = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!roomName.trim() || !roomDescription.trim()) {
      toast("Room name and description are required");
      return;
    }
    await createRoom({ name: roomName, description: roomDescription });
    if (success) {
      setIsModalOpen(false);
      setRoomName("");
      setRoomDescription("");
    }
  };

  const handleJoinByCode = async (e: React.MouseEvent) => {
    e.stopPropagation();
    if (!roomCode.trim()) {
      toast("Room code required");
      return;
    }

    await joinRoom({ room_code: roomCode });
    getRooms();
  };

  const handleDeleteRoom = async (room_code: string, e: React.MouseEvent) => {
    e.stopPropagation();
    await deleteRoom({ room_code });
    if (success) {
      toast("Room removed successfully");
    }
  };

  const handleSettings = () => {
    router.push("/settings");
  };

  const handleLogout = () => {
    router.push("/login");
  };

  const handleLeaveRoom = async (room_code: string, e: React.MouseEvent) => {
    e.stopPropagation();

    await leaveRoom({ room_code });

    if (success) {
      toast("Leave room successful");
    }
  };

  if (loading || !user) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="backdrop-blur-xl border border-white/10 p-6 rounded-xl">
          <p className="text-white/80">Loading your chat rooms...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen flex flex-col items-center p-4 py-12">
      <div className="w-full max-w-3xl backdrop-blur-xl border border-white/10 p-6 rounded-2xl animate-fadeIn">
        <div className="space-y-6">
          <div className="flex items-center justify-between">
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <div className="flex items-center gap-3 cursor-pointer hover:bg-white/10 duration-200 p-2 rounded-lg transition-colors">
                  {user?.profile_url ? (
                    <Image
                      width={40}
                      height={40}
                      src={user?.profile_url}
                      alt="Profile"
                      className="w-10 h-10 rounded-full object-cover"
                    />
                  ) : (
                    <div className="w-10 h-10 rounded-full bg-white/10 flex items-center justify-center">
                      <User size={20} className="text-white/70" />
                    </div>
                  )}
                  <span className="font-medium text-white">{user?.username}</span>
                </div>
              </DropdownMenuTrigger>
              <DropdownMenuContent className="bg-black border border-white/10 text-white">
                <DropdownMenuItem
                  onClick={handleSettings}
                  className="flex items-center gap-2 cursor-pointer hover:bg-white/10"
                >
                  <Settings size={16} />
                  Settings
                </DropdownMenuItem>
                <DropdownMenuItem
                  onClick={handleLogout}
                  className="flex items-center gap-2 cursor-pointer hover:bg-white/10 text-red-400 hover:text-red-500"
                >
                  <LogOut size={16} />
                  Logout
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </div>

          <div className="flex items-center justify-between">
            <h1 className="text-2xl font-bold tracking-tighter">
              Your Chat Rooms
            </h1>
            <Button
              variant="outline"
              size="sm"
              onClick={() => setIsModalOpen(true)}
              className="flex items-center gap-2 hover:opacity-70 cursor-pointer"
            >
              <Plus size={16} />
              New Room
            </Button>
          </div>
          <div className="backdrop-blur-xl border border-white/10 p-4 rounded-lg">
            <h2 className="font-medium mb-3">Join a room by code</h2>
            <div className="flex gap-2">
              <Input
                type="text"
                placeholder="Enter room code"
                className="bg-white/5 border-white/10 focus:border-white/30"
                value={roomCode}
                onChange={(e) => setRoomCode(e.target.value)}
                autoComplete="off"
              />
              <Button
                onClick={handleJoinByCode}
                className="flex items-center gap-1 bg-white text-black font-semibold hover:opacity-70 cursor-pointer whitespace-nowrap"
              >
                Join
                <ArrowRight size={16} />
              </Button>
            </div>
          </div>

          <ScrollArea className="h-[50vh]">
            <div className="space-y-3 overflow-y-auto">
              {rooms.length > 0 ? (
                rooms.map((room) => (
                  <div
                    key={room.id}
                    onClick={() => handleJoinRoom(room.room_code)}
                    className="backdrop-blur-xl border border-white/10 p-4 rounded-lg cursor-pointer hover:bg-white/10 transition-colors flex items-center justify-between group"
                  >
                    <div className="flex items-center gap-3">
                      <div className="w-10 h-10 rounded-full bg-white/10 flex items-center justify-center">
                        <MessageSquare size={20} className="text-white/70" />
                      </div>
                      <div>
                        <h3 className="font-medium">{room.name || room.id}</h3>
                        <div className="flex items-center text-xs text-white/60 gap-2">
                          <Clock size={12} />
                          <span>{formatDate(new Date(room.created_at))}</span>
                        </div>
                      </div>
                    </div>
                    {user.id == Number(room.owner_id) ? (
                      <Button
                        variant="ghost"
                        size="sm"
                        onClick={(e) => handleDeleteRoom(room.room_code, e)}
                        className="opacity-0 group-hover:opacity-100 transition-opacity cursor-pointer hover:opacity-70"
                      >
                        <Trash2 size={18} className="text-white/70" />
                      </Button>
                    ) : (
                      <Button
                        variant="ghost"
                        size="sm"
                        onClick={(e) => handleLeaveRoom(room.room_code, e)}
                        className="opacity-0 group-hover:opacity-100 transition-opacity cursor-pointer hover:opacity-70"
                      >
                        <DoorOpen size={18} className="text-white/70" />
                      </Button>
                    )}
                  </div>
                ))
              ) : (
                <div className="text-center p-8 text-white/60">
                  <p>No chat rooms found</p>
                  <p className="text-sm mt-2">
                    Create a new room to get started
                  </p>
                </div>
              )}
            </div>
          </ScrollArea>

          <div className="text-center text-xs text-white/40">
            Click on a room to join the conversation
          </div>
        </div>
      </div>

      <Dialog open={isModalOpen} onOpenChange={setIsModalOpen}>
        <DialogContent className="sm:max-w-[425px] backdrop-blur-xl border border-white/10 bg-black/80 text-white">
          <DialogHeader>
            <DialogTitle>Create New Room</DialogTitle>
          </DialogHeader>
          <form onSubmit={handleCreateRoom}>
            <div className="grid gap-4 py-4">
              <div className="grid gap-2">
                <label htmlFor="roomName" className="text-sm font-medium">
                  Room Name
                </label>
                <Input
                  id="roomName"
                  value={roomName}
                  onChange={(e) => setRoomName(e.target.value)}
                  placeholder="Enter room name"
                  className="bg-white/5 border-white/10 focus:border-white/30"
                  autoComplete="off"
                  required
                />
              </div>
              <div className="grid gap-2">
                <label htmlFor="roomDescription" className="text-sm font-medium">
                  Description
                </label>
                <Textarea
                  id="roomDescription"
                  value={roomDescription}
                  onChange={(e) => setRoomDescription(e.target.value)}
                  placeholder="Enter room description (optional)"
                  className="bg-white/5 border-white/10 focus:border-white/30"
                />
              </div>
            </div>
            <DialogFooter>
              <Button
                type="button"
                variant="outline"
                onClick={() => setIsModalOpen(false)}
                className="border-white/10 text-white hover:bg-white/10"
              >
                Cancel
              </Button>
              <Button
                type="submit"
                className="bg-white text-black hover:opacity-70"
              >
                Create Room
              </Button>
            </DialogFooter>
          </form>
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default SelectChats;