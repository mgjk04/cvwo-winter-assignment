"use  client";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Stack from "@mui/material/Stack";
import TextField from "@mui/material/TextField";
import FormHelperText from "@mui/material/FormHelperText";
import { z } from "zod";
import { FieldError, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { commentFormSchema } from "@/app/(client)/comments/zod";
import useCreateComment from "@/app/(client)/comments/_hooks/useCreateComment";
import { handleError } from "@/app/(client)/comments/utils";

export default function CreateComment(ModifyFormProps: {
  submitURL: string;
  comment?: z.infer<typeof commentFormSchema>;
}) {
  const { mutate } = useCreateComment(ModifyFormProps.submitURL);

  async function onSubmit(values: z.infer<typeof commentFormSchema>) {
    mutate(values, {
      onError: (error) => {
        handleError(setError, error, () =>
          mutate(values, {
            onError: (error) => {
              console.error(error);
              setError("root.auth", {
                message: "Login to perform this action",
              });
            },
          }),
        );
      },
      onSuccess: () => {
        reset();
      },
    });
  }

  const {
    register,
    formState: { errors, isSubmitting, isDirty },
    setError,
    handleSubmit,
    reset,
  } = useForm({
    resolver: zodResolver(commentFormSchema),
    defaultValues: {
      content: ModifyFormProps.comment?.content || "",
    },
  });

  return (
    <Box>
      <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
        <Stack className="flex w-full">
          <TextField
            {...register("content")}
            fullWidth
            id="content"
            name="content"
            label="Join the conversation"
            variant="outlined"
            required
            placeholder="Share your thoughts"
            error={!!errors.content || !!errors.root}
            helperText={errors.content?.message}
          />
          <FormHelperText error={!!errors.root}>
            {(Object.values(errors.root || {}) as FieldError[])[0]?.message}
          </FormHelperText>
          {isDirty && (
            <Button
              size="small"
              variant="contained"
              type="submit"
              disabled={isSubmitting}
            >
              Create!
            </Button>
          )}
        </Stack>
      </form>
    </Box>
  );
}
