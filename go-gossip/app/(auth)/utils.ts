import { HttpStatus } from "http-status-ts";
import { UseFormSetError } from "react-hook-form";
import z from "zod";
import { userCredentialsSchema } from "./zod";

//no retry on auth failure
export function handleError(
  setError: UseFormSetError<z.infer<typeof userCredentialsSchema>>,
  error: Error,
) {
  switch (parseInt(error.message)) {
    case HttpStatus.NOT_FOUND:
      setError("root.auth", { message: "Invalid Credentials" }); //well, in the auth context its invalid login credentials
      break;
    case HttpStatus.FORBIDDEN:
      setError("root.auth", { message: "You are unauthorized" });
      break;
    case HttpStatus.UNAUTHORIZED:
      setError("root.auth", { message: "Login to perform this action" });
      break;
    case HttpStatus.CONFLICT:
      setError("username", {
        message: "This username has beem used! Try another!",
      });
      break;
    case HttpStatus.BAD_REQUEST:
      setError("root.invalid", { message: "Please, no funny business" }); //change this, message later
      break;
    default:
      setError("root.server", {
        message: "Something went wrong, try again later",
      });
  }
}
