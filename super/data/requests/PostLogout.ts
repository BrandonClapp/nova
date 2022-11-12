import { HttpDefinition } from "../HttpDefinition";

export const PostLogout: HttpDefinition<
  undefined,
  { success: string; message: string }
> = {
  method: "POST",
  endpoint: "/auth/logout",
};
