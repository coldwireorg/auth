<script>
  export let username
</script>

<span>{username.slice(0, 1).toUpperCase()}</span>

<style>
  span {
    width: 48px;
    height: 48px;
    border-radius: 50%;
    background-color: var(--secondary-violet);

    font-size: 18px;

    display: flex;
    justify-content: center;
    align-items: center;

    user-select: none;
  }
</style>