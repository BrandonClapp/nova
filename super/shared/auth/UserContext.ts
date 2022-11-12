// How do I distribute 2 versions of react in the same repository
import { createContext } from "react";

export const UserContext = createContext({
  user: undefined,
  loading: true,
  refetch: (body?: any) => {}, // Pretty sure react hates this
});
