import z from "zod";

export const topicFormSchema = z.object({
  topicname: z.string(),
  description: z.string().optional(),
});