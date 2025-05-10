import { cn } from '@/libs/utils';
import { Member, Room } from '@/types/rooms.type';

interface UserListProps {
  users: Member[];
  room: Room;
}

const UserList = ({ users, room }: UserListProps) => {
  return (
    <div className="p-4 animate-fadeIn">
      <h2 className="font-medium mb-4 text-white/80">People in this room</h2>
      <div className="space-y-2">
        {users.map((user) => (
          <div 
            key={user.user_id} 
            className="flex items-center space-x-2 p-2 rounded-lg hover:bg-white/5"
          >
            <div className="relative">
              <div className="w-8 h-8 rounded-full bg-white/10 flex items-center justify-center">
                {user.username.charAt(0).toUpperCase()}
              </div>
              <div 
                className={cn(
                  "absolute bottom-0 right-0 w-2.5 h-2.5 rounded-full border-2 border-black",
                  user.online ? "bg-green-400" : "bg-gray-400"
                )}
              />
            </div>
            <div className="flex-1 min-w-0">
              <p className="text-sm truncate">{user.username} {user.user_id === room.owner_id ? "👑" : "🥷🏼"}</p>
              <p className="text-xs text-white/50">
                {user.online ? "Online" : "Offline"}
              </p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default UserList;
