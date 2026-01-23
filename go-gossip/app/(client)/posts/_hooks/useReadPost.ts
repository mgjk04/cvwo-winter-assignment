import { useQuery } from "@tanstack/react-query";

const readPost = async (readURL: string) => {
    const res = await fetch(readURL, {
        method: "GET",
        headers: { "Content-Type": "application/json" },
        credentials: "include"
    });
    if(!res.ok) throw new Error(String(res.status));
    return res.json();
}

export default function useReadPost(readURL: string) {
  return useQuery({
    queryKey: ['post', readURL],
    queryFn: () => readPost(readURL),
  });
}