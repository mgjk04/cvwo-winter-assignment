import z from "zod";
import { commentFormSchema } from "../zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";

const editComment = (editURL: string) => async (values: z.infer<typeof commentFormSchema>) => {
    const res = await fetch(editURL, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(values),
        credentials: "include"
    });
    if(!res.ok) throw new Error(String(res.status));
    return res.json();
}

export default function useEditComment(editURL: string){
    const client = useQueryClient();
    return useMutation({
    mutationFn: editComment(editURL),
    onSettled: () => {
      client.invalidateQueries({ queryKey: ["comment"] });
    },
  });
}