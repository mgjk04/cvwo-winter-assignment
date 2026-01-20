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
import { topicFormSchema } from "../../../zod";
import { handleError } from "../../../utils";
import useEditTopic from "@/app/topics/_hooks/useEditTopic";

export default function EditTopic(ModifyFormProps: {
  submitURL: string;
  topic?: z.infer<typeof topicFormSchema>;
}) {
  const { mutate } = useEditTopic(ModifyFormProps.submitURL);

  async function onSubmit(values: z.infer<typeof topicFormSchema>) {
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
    });
  }

  const {
    register,
    formState: { errors, isSubmitting },
    setError,
    handleSubmit,
    watch
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
            Edit Topic
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
                defaultValue={ModifyFormProps.topic?.topicname}
                slotProps={{
                  // eslint-disable-next-line react-hooks/incompatible-library
                  inputLabel: { shrink: !!watch('topicname') },
                }}
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
                defaultValue={ModifyFormProps.topic?.description}
                slotProps={{
                  inputLabel: { shrink: !!watch('description') },
                }}
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
