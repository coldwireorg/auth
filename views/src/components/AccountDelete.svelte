<script>
  import Card from '../components/Card.svelte'
  import Text from '../components/Text.svelte'
  import Field from '../components/Field.svelte'
  import Button from '../components/Button.svelte'

  const del = () => [
    fetch("/api/user/", {
      method: 'DELETE'
    })
    .then(() => window.location.reload())
  ]
</script>

<Card style="display: flex; flex-direction: column; gap:8px;">
  <Text type="h3">
    Delete your account
  </Text>
  <br>
  <Field type="password" placeholder="Enter password" />
  <br>
  <Button color="red">
    Delete account
  </Button>
</Card>
