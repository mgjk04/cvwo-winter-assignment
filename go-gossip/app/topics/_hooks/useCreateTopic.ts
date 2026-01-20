import { useMutation, useQueryClient } from "@tanstack/react-query";
import z from "zod";
import { topicFormSchema } from "../zod";
import "dotenv/config";

const createURL = process.env.API_URL + "/topics/";

const createTopic = async (values: z.infer<typeof topicFormSchema>) => {
  const res = await fetch(createURL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(values),
    credentials: "include"
  });
  if (!res.ok) throw new Error(String(res.status));
  return res.json();
};

export default function useCreateTopic() {
  const client = useQueryClient();
  return useMutation({
    mutationFn: createTopic,
    onSettled: () => {
      client.invalidateQueries({ queryKey: ["topic"] });
    },
  });
}
