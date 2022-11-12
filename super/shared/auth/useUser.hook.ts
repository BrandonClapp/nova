import { CurrentUser } from "@/models/CurrentUser";
import { useContext } from "react";
import { UserContext } from "./UserContext";

export interface UserState {
  user: CurrentUser | undefined;
  loading: boolean;
  refetch: (body?: any) => void;
}

export const useUser = () => {
  const userState = useContext<UserState>(UserContext);
  return userState;
};
