export default function ItemList(props: { items: JSX.Element[] }) {
  return (
    <div className="flex flex-col gap-1">
      {props?.items?.map((item) => item)}
    </div>
  );
}
