import { Colors } from "../shared/types/TailwindColors";

export interface Payment {
  id: string;
  name: string;
  date: string;
  amount: number;
  category: {
    name: string;
    color: Colors;
  };
  isDeposit: boolean;
  balance: number;
  status: "paid" | "pending";
}
