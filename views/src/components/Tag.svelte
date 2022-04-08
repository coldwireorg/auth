<script>
  export let text
  export let color
  export let background

  if (background == "purple") {
    background = 'rgba(204, 49, 124, 0.25)'
    color = '#CC317C'
  }

  if (background == "green") {
    background = 'var(--primary-green-25)'
    color = 'var(--primary-green)'
  }
</script>

<span style="background-color: {background}; color: {color}">{text}</span>

<style>
  span {
    border-radius: 8px;

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 0px 8px;
  }
</style>