import z from "zod";

export const commentFormSchema = z.object({
  content: z.string(),
});