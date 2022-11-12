import classNames from "classnames";
import { Colors } from "./types/TailwindColors";

export default function Badge(props: { children: string; color: Colors }) {
  return (
    <span
      className={classNames(
        "inline-flex items-center rounded-md bg-gray-100 px-3 py-0.5 text-sm font-medium text-gray-800",
        { "bg-gray-100 text-gray-800": props.color === Colors.gray },
        { "bg-green-100 text-green-800": props.color === Colors.green },
        { "bg-red-100 text-red-800": props.color === Colors.red },
        { "bg-yellow-100 text-yellow-800": props.color === Colors.yellow },
        { "bg-blue-100 text-blue-800": props.color === Colors.blue },
        { "bg-indigo-100 text-indigo-800": props.color === Colors.indigo },
        { "bg-purple-100 text-purple-800": props.color === Colors.purple },
        { "bg-pink-100 text-pink-800": props.color === Colors.pink }
      )}
    >
      {props.children}
    </span>
  );
}
