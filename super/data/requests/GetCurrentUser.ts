import { CurrentUser } from "../../models/CurrentUser";
import { HttpDefinition } from "../HttpDefinition";

export const GetCurrentUser: HttpDefinition<undefined, CurrentUser> = {
  method: "GET",
  endpoint: "/auth/current-user",
};
