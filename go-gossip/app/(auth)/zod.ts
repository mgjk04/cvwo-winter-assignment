import z from 'zod'
const alphanumerRegex = new RegExp(/^[\w0-9]+$/); //include underscorde
export const userCredentialsSchema = z.object({
    username: z.string().trim().regex(alphanumerRegex, "Only letters, numbers and _ allowed")
})