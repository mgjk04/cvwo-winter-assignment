import z from "zod";
import { postFormSchema } from "../zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";



const editURL = 'http://localhost:8080/posts';

const editPost = async (values: z.infer<typeof postFormSchema>) => {
    const res = await fetch(editURL, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(values),
        credentials: "include"
    });
    if(!res.ok) throw new Error(String(res.status));
    return res.json();
}

export default function useEditPost(){
    const client = useQueryClient();
    return useMutation({
    mutationFn: editPost,
    onSettled: () => {
      client.invalidateQueries({ queryKey: ["post"] });
    },
  });
}