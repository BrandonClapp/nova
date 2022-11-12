import { CurrentUser } from "../../models/CurrentUser";
import { HttpDefinition } from "../HttpDefinition";

export const PostLogin: HttpDefinition<
  { email: string; password: string },
  CurrentUser
> = {
  method: "POST",
  endpoint: "/auth/login",
};
