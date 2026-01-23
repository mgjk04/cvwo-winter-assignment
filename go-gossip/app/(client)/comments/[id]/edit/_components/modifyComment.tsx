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
import { commentFormSchema } from "@/app/(client)/comments/zod";
import { handleError } from "@/app/(client)/comments/utils";
import useCreateComment from "@/app/(client)/comments/_hooks/useCreateComment";

export default function CreateComment(ModifyFormProps: {
  submitURL: string;
  comment?: z.infer<typeof commentFormSchema>;
}) {
  const { mutate } = useCreateComment(ModifyFormProps.submitURL);

  async function onSubmit(values: z.infer<typeof commentFormSchema>) {
    mutate(values, {
      onError: (error) => {
        handleError(setError, error, () => mutate(values));
      },
    });
  }

  const {
    register,
    formState: { errors, isSubmitting },
    setError,
    handleSubmit,
  } = useForm({
    resolver: zodResolver(commentFormSchema),
    defaultValues: {
      content: ModifyFormProps.comment?.content || "",
    },
  });

  return (
    <Box>
      <Card>
        <CardContent>
          <Typography className="strong" variant="h4">
            Edit Comment
          </Typography>
          <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
            <Stack className="flex w-full gap-2.5">
              <TextField
                {...register("content")}
                fullWidth
                id="content"
                name="content"
                label="Content"
                variant="outlined"
                required
                placeholder="Share your thoughts"
                error={!!errors.content || !!errors.root}
                helperText={errors.content?.message}
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
                  Edit!
                </Button>
              </CardActions>
            </Stack>
          </form>
        </CardContent>
      </Card>
    </Box>
  );
}
