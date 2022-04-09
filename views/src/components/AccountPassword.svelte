<script>
  import Card from '../components/Card.svelte'
  import Text from '../components/Text.svelte'
  import Field from '../components/Field.svelte';
  import Button from '../components/Button.svelte';

  let password = ""
  let newpassword = ""
  let newrepassword = ""
  
  let error = ""
  let success = ""

  const change = () => {
    error = ""
    if (newpassword != "" && password != "") {
        if (newpassword != newrepassword) {
        error = "Password does not match"
        return
      }

      fetch("/api/user/password", {
        method: 'PUT',
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          password: password,
          new: newpassword
        })
      })
      .then(req => req.json())
      .then(data => {
        if (data.status == "SUCCESS") {
          success = "Password updated!"
        } else {
          error = data.message
        }
      })
    } else {
      error = "Please fill the fields"
    }
  }
</script>

<Card style="display: flex; flex-direction: column; gap:8px;">
  <Text type="h3">
    Change your password
  </Text>
  <br>
  {#if error != ""}
    <span class="error">{error}</span>
  {/if}
  {#if success != ""}
    <span class="success">{success}</span>
  {/if}
  <Field type="password" placeholder="Enter your new password" bind:value={newpassword} />
  <Field type="password" placeholder="Repeat your new password" bind:value={newrepassword} />
  <Field type="password" placeholder="Enter your current password" bind:value={password} />
  <br>
  <Button on:click={change}>
    Change password
  </Button>
</Card>

<style>
  .error {
    background-color: var(--secondary-red);
    padding: 8px;
    border-radius: 8px;
  }

  .success {
    background-color: var(--secondary-green);
    padding: 8px;
    border-radius: 8px;
  }
</style>