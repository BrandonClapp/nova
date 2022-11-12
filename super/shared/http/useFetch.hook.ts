import { useState } from "react";
import { HttpDefinition } from "./HttpDefinition";
import { HttpState } from "./HttpState";
import { performRequest } from "./performRequest";

const HOST = process.env.NEXT_PUBLIC_API_HOST || "https://localhost:8080";

export function useFetch<TBody, TReturn>(
  props: HttpDefinition<TBody, TReturn>
) {
  const { method, endpoint } = props;

  const [data, setData] = useState<TReturn>();
  const [error, setError] = useState(undefined);
  const [state, setState] = useState<HttpState>(HttpState.NotStarted);

  const request = (body?: TBody) => {
    setData(undefined);
    setError(undefined);
    setState(HttpState.Pending);

    performRequest({ method, url: `${HOST}${endpoint}`, data: body })
      .then(
        (success) => {
          setData(success as TReturn);
          setError(undefined);
        },
        (rejected) => {
          setData(undefined);
          setError(rejected);
        }
      )
      .finally(() => {
        setState(HttpState.Complete);
      });
  };

  return { request, data, error, state };
}
