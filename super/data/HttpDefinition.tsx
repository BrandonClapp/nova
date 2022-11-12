export default () => {};

export interface HttpDefinition<TBody, TReturn> {
  method: "GET" | "POST";
  endpoint: string;
}
