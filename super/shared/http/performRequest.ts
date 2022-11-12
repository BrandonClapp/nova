function getFetchOptions(method: "GET" | "POST", data?: any): RequestInit {
  // Default options are marked with *
  return {
    method: method, // *GET, POST, PUT, DELETE, etc.
    mode: "cors", // no-cors, *cors, same-origin
    cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
    credentials: "include", // include, *same-origin, omit
    headers: {
      "Content-Type": "application/json",
      // 'Content-Type': 'application/x-www-form-urlencoded',
    },
    redirect: "follow", // manual, *follow, error
    referrerPolicy: "no-referrer", // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
    body: data ? JSON.stringify(data) : undefined, // body data type must match "Content-Type" header
  };
}

export async function performRequest<T>({
  method,
  url = "",
  data,
}: {
  method: "GET" | "POST";
  url: string;
  data: any;
}): Promise<T> {
  return new Promise(async (resolve, reject) => {
    const fetchOptions = getFetchOptions(method, data);
    const response = await fetch(url, fetchOptions);

    const body = await response.json();

    if (!response.ok) {
      reject(body.error || "An error occured while processing the request.");
    }

    resolve(body);
  });
}
