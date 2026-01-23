import z from "zod";

export const postFormSchema = z.object({
  title: z.string(),
  description: z.string().optional(),
});