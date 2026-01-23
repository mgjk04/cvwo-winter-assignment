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

export default function useDeleteComment(){
    const client = useQueryClient();
    return useMutation({
    mutationFn: (deleteURL: string) => deleteComment(deleteURL),
    onSettled: () => {
      client.invalidateQueries({ queryKey: ["comment"] });
    },
  });
}