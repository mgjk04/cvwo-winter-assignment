import z from "zod";
import { topicFormSchema } from "../zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";



const editURL = 'http://localhost:8080/posts';

const editTopic = async (values: z.infer<typeof topicFormSchema>) => {
    const res = await fetch(editURL, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(values),
        credentials: "include"
    });
    if(!res.ok) throw new Error(String(res.status));
    return res.json();
}

export default function useEditTopic(){
    const client = useQueryClient();
    return useMutation({
    mutationFn: editTopic,
    onSettled: () => {
      client.invalidateQueries({ queryKey: ["topic"] });
    },
  });
}