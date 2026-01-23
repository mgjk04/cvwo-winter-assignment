import z from "zod";
import { postFormSchema } from "../zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";

const editPost = (editURL: string) => async (values: z.infer<typeof postFormSchema>) => {
    const res = await fetch(editURL, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(values),
        credentials: "include"
    });
    if(!res.ok) throw new Error(String(res.status));
}

export default function useEditPost(editURL: string){
    const client = useQueryClient();
    return useMutation({
    mutationFn: editPost(editURL),
    onSettled: () => {
      client.invalidateQueries({ queryKey: ["post"] });
    },
  });
}