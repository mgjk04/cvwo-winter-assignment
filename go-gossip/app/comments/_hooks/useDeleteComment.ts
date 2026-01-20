import { useMutation, useQueryClient } from "@tanstack/react-query";

const deleteComment = async (deleteURL: string) => {
    const res = await fetch(deleteURL, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
        credentials: "include"
    });
    if(!res.ok) throw new Error(String(res.status));
    return res.json();
}

export default function useEditComment(deleteURL: string){
    const client = useQueryClient();
    return useMutation({
    mutationFn: () => deleteComment(deleteURL),
    onSettled: () => {
      client.invalidateQueries({ queryKey: ["post"] });
    },
  });
}