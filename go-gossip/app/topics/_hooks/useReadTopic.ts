import { useQuery } from "@tanstack/react-query";

const readTopic = async (readURL: string) => {
    const res = await fetch(readURL, {
        method: "GET",
        headers: { "Content-Type": "application/json" },
        credentials: "include"
    });
    if(!res.ok) throw new Error(String(res.status));
    return res.json();
}

export default function useReadTopic(readURL: string) {
  return useQuery({
    queryKey: ['topic'],
    queryFn: () => readTopic(readURL),
  });
}