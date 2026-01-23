import z from "zod";
import { topicFormSchema } from "../zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";

const editTopic = (editURL: string) => async (values: z.infer<typeof topicFormSchema>) => {
    const res = await fetch(editURL, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(values),
        credentials: "include"
    });
    if(!res.ok) throw new Error(String(res.status));
}

export default function useEditTopic(editURL: string){
    const client = useQueryClient();
    return useMutation({
    mutationFn: editTopic(editURL),
    onSettled: () => {
      client.invalidateQueries({ queryKey: ["topic"] });
    },
  });
}