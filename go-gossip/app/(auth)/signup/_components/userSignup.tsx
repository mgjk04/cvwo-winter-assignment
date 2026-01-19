"use  client";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Stack from "@mui/material/Stack";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import FormHelperText from "@mui/material/FormHelperText";
import { z } from "zod";
import { FieldError, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardActions from "@mui/material/CardActions";
import { handleError } from "../../utils";
import { userCredentialsSchema } from "../../zod";
import useSignup from "../_hooks/useSignup";

export default function UserCredentials() {
  const { mutate } = useSignup();
  async function onSubmit(values: z.infer<typeof userCredentialsSchema>) {
    mutate(values, {
      onError: (error) => {
        handleError(setError, error);
      },
    });
  }

  const {
    register,
    formState: { errors, isSubmitting },
    setError,
    handleSubmit,
  } = useForm({
    resolver: zodResolver(userCredentialsSchema),
    defaultValues: {
      username: "",
    },
  });

  return (
    <Box>
      <Card>
        <CardContent>
          <Typography className="strong" variant="h4">
            Sign Up
          </Typography>
          <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
            <Stack className="flex w-full gap-2.5">
              <TextField
                {...register("username")}
                fullWidth
                id="username"
                name="username"
                label="Username"
                variant="outlined"
                required
                placeholder="What is your name?"
                error={!!errors.username || !!errors.root}
                helperText={errors.username?.message}
              />
              <FormHelperText error={!!errors.root}>
                {(Object.values(errors.root || {}) as FieldError[])[0]?.message}
              </FormHelperText>
              <CardActions>
                <Button
                  variant="contained"
                  type="submit"
                  disabled={isSubmitting}
                >
                  Sign Up!
                </Button>
              </CardActions>
            </Stack>
          </form>
        </CardContent>
      </Card>
    </Box>
  );
}
