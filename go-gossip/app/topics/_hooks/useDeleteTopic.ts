import { useMutation, useQueryClient } from "@tanstack/react-query";

const deleteTopic = async (deleteURL: string) => {
    const res = await fetch(deleteURL, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
        credentials: "include"
    });
    if(!res.ok) throw new Error(String(res.status));
    return res.json();
}

export default function useEditTopic(deleteURL: string){
    const client = useQueryClient();
    return useMutation({
    mutationFn: () => deleteTopic(deleteURL),
    onSettled: () => {
      client.invalidateQueries({ queryKey: ["topic"] });
    },
  });
}