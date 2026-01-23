import { HttpStatus } from "http-status-ts";
import { UseFormSetError } from "react-hook-form";
import { topicFormSchema } from "./zod";
import z from "zod";
import { refresh } from "../(auth)/refresh";
import { setCookie } from "cookies-next";

//ngl looks pretty messy
export async function handleError(setError: UseFormSetError<z.infer<typeof topicFormSchema>>, error: Error, retry: () => void){
    console.error(error);
    switch(parseInt(error.message)) {
        case HttpStatus.FORBIDDEN:
            setError('root.auth', { message: "You are unauthorized"});
            break;
        case HttpStatus.UNAUTHORIZED:
            try {
                const res: {user_id: string} = await refresh();
                retry();
                setCookie("user_id", res.user_id);
            } catch (error) {
                console.log(error);
                setError('root.auth', { message: "Login to perform this action"});
            }
            break;
        case HttpStatus.CONFLICT:
            setError('topicname', { message: "This title has beem used! Try another!"});
            break;
        case HttpStatus.BAD_REQUEST:
            setError('root.invalid', {message: "Please, no funny business"}); //change this, message later 
            break;
        default:
            setError('root.server', {message: "Something went wrong, try again later"});

    }
}