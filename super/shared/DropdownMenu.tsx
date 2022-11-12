import { Menu, Transition } from "@headlessui/react";
import classNames from "classnames";
import { Fragment } from "react";

export default function DropdownMenu(props: {
  children: JSX.Element;
  links: { href?: string; name: string; onClick?: () => void }[];
}) {
  return (
    <Menu as="div" className="relative ml-3">
      <div>
        <Menu.Button className="flex max-w-xs items-center rounded-full bg-gray-100 text-sm focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800">
          <span className="sr-only">Open menu</span>
          {props.children}
        </Menu.Button>
      </div>
      <Transition
        as={Fragment}
        enter="transition ease-out duration-100"
        enterFrom="transform opacity-0 scale-95"
        enterTo="transform opacity-100 scale-100"
        leave="transition ease-in duration-75"
        leaveFrom="transform opacity-100 scale-100"
        leaveTo="transform opacity-0 scale-95"
      >
        <Menu.Items className="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
          {props.links.map((item) => (
            <Menu.Item key={item.name}>
              {({ active }) => (
                <a
                  href={item.href}
                  onClick={item.onClick}
                  className={classNames(
                    active ? "bg-gray-100" : "",
                    "cursor-pointer block px-4 py-2 text-sm text-gray-700"
                  )}
                >
                  {item.name}
                </a>
              )}
            </Menu.Item>
          ))}
        </Menu.Items>
      </Transition>
    </Menu>
  );
}