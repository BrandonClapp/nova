import { GetCurrentUser } from "@/data/requests/GetCurrentUser";
import { CurrentUser } from "@/models/CurrentUser";
import React, { useEffect, useState } from "react";
import { HttpState } from "../http/HttpState";
import { useFetch } from "../http/useFetch.hook";
import { UserContext } from "./UserContext";
import { UserState } from "./useUser.hook";

export default function UserContextProvider(props: { children: JSX.Element }) {
  const { children } = props;

  const { request, data, error, state } = useFetch<undefined, CurrentUser>(
    GetCurrentUser
  );

  const [userState, setUserState] = useState<UserState>({
    user: data,
    loading: true,
    refetch: () => {
      request();
    },
  });

  useEffect(() => {
    request();
  }, []);

  useEffect(() => {
    let isLoading = false;
    if (state === HttpState.NotStarted || state === HttpState.Pending) {
      isLoading = true;
    }

    if (state === HttpState.Complete) {
      isLoading = false;
    }

    setUserState({
      ...userState,
      user: data,
      loading: isLoading,
    });
  }, [data, state, error]);

  return (
    <UserContext.Provider value={userState}>{children}</UserContext.Provider>
  );
}
