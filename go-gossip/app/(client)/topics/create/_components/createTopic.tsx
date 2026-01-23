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
import { topicFormSchema } from "../../zod";
import useCreateTopic from "../../_hooks/useCreateTopic";
import { handleError } from "../../utils";
import { useRouter } from "next/navigation";

export default function CreateTopic(ModifyFormProps: {
  submitURL: string;
  topic?: z.infer<typeof topicFormSchema>;
}) {
  const router = useRouter();
  const { mutate } = useCreateTopic(router);

  async function onSubmit(values: z.infer<typeof topicFormSchema>) {
    mutate(values, {
      onError: (error) => {
        handleError(setError, error, () =>
          mutate(values, {
            onSuccess: () => {
              router.push("/topics");
            },
          }),
        );
      },
      onSuccess: () => {
        router.push("/topics");
      },
    });
  }

  const {
    register,
    formState: { errors, isSubmitting },
    setError,
    handleSubmit,
  } = useForm({
    resolver: zodResolver(topicFormSchema),
    defaultValues: {
      topicname: ModifyFormProps.topic?.topicname,
      description: ModifyFormProps.topic?.description,
    },
  });

  return (
    <Box>
      <Card>
        <CardContent>
          <Typography className="strong" variant="h4">
            Create Topic
          </Typography>
          <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
            <Stack className="flex w-full gap-2.5">
              <TextField
                {...register("topicname")}
                fullWidth
                id="title"
                name="topicname"
                label="Title"
                variant="outlined"
                required
                placeholder="What to call the topic?"
                error={!!errors.topicname || !!errors.root}
                helperText={errors.topicname?.message}
              />
              <TextField
                {...register("description")}
                fullWidth
                multiline
                id="description"
                name="description"
                label="Description"
                variant="outlined"
                placeholder="What should others know about this?"
                error={!!errors.description || !!errors.root}
                helperText={errors.description?.message}
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
                  Create!
                </Button>
              </CardActions>
            </Stack>
          </form>
        </CardContent>
      </Card>
    </Box>
  );
}
