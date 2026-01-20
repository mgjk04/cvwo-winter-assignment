import { HttpStatus } from "http-status-ts";
import { UseFormSetError } from "react-hook-form";
import { commentFormSchema } from "./zod";
import z from "zod";
import { refresh } from "../(auth)/refresh";

//ngl looks pretty messy
export function handleError(setError: UseFormSetError<z.infer<typeof commentFormSchema>>, error: Error, retry: () => void){
    console.error(error);
    switch(parseInt(error.message)) {
        case HttpStatus.FORBIDDEN:
            setError('root.auth', { message: "You are unauthorized"});
            break;
        case HttpStatus.UNAUTHORIZED:
            try {
                refresh();
                retry();
            } catch (error) {
                console.log(error);
                setError('root.auth', { message: "Login to perform this action"});
            }
            break;
        case HttpStatus.BAD_REQUEST:
            setError('root.invalid', {message: "Please, no funny business"}); //change this, message later 
            break;
        default:
            setError('root.server', {message: "Something went wrong, try again later"});

    }
}