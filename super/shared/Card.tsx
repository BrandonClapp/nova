export default function Card(props: { children: JSX.Element }) {
  return (
    <div className="p-4 border bg-white rounded-md shadow-sm">
      {props.children}
    </div>
  );
}
