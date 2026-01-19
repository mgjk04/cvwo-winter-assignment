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
import { verb } from "../types";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardActions from "@mui/material/CardActions";
import { postFormSchema } from "../zod";
import { handleError } from "../utils"
import useCreatePost from "../_hooks/useCreatePost";



export default function ModifyPost(ModifyFormProps: {
  verb: verb;
  submitURL: string;
  post?: z.infer<typeof postFormSchema>;
}) {
  const { mutate } = useCreatePost();

  async function onSubmit(values: z.infer<typeof postFormSchema>) {
    mutate(values, {
      onError: (error) => {
        handleError(setError, error, () => mutate(values));
      }
    })
  }

  const {
    register,
    formState: { errors, isSubmitting },
    setError,
    handleSubmit,
  } = useForm({
    resolver: zodResolver(postFormSchema),
    defaultValues: {
      title: ModifyFormProps.post?.title,
      description: ModifyFormProps.post?.description,
    },
  });

  return (
    <Box>
      <Card>
        <CardContent>
          <Typography className='strong' variant="h4">{ModifyFormProps.verb} Post</Typography>
          <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
            <Stack className="flex w-full gap-2.5">
              <TextField
                {...register("title")}
                fullWidth
                id="title"
                name="topicname"
                label="Title"
                variant="outlined"
                required
                placeholder="Share your thoughts"
                error={!!errors.title || !!errors.root}
                helperText={errors.title?.message}
              />
              <TextField
                {...register("description")}
                fullWidth
                multiline
                id="description"
                name="description"
                label="Description"
                variant="outlined"
                placeholder="Details?"
                error={!!errors.description || !!errors.root}
                helperText={errors.description?.message}
              />
              <FormHelperText error={!!errors.root}>{(Object.values(errors.root || {}) as FieldError[])[0]?.message}</FormHelperText>
              <CardActions>
                <Button
                  variant="contained"
                  type="submit"
                  disabled={isSubmitting}
                >
                  {ModifyFormProps.verb}!
                </Button>
              </CardActions>
            </Stack>
          </form>
        </CardContent>
      </Card>
    </Box>
  );
}
