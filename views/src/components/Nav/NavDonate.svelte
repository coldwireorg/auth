<script>
  import Card from '../Card.svelte'
  import Text from '../Text.svelte'
  import Button from '../Button.svelte';
</script>

<Card type="green-25" style="display: flex; flex-direction: column; gap:16px">
  <Text style="font-size:14px;" type="h3">Help us to keep our services alive!</Text>
  <div class="about">
    <Button color="green" link="https://coldwire.org/about#donate">
      Donate
      <img src="/icons/heart.svg" alt="">
    </Button>
    <a href="https://coldwire.org/about#donate" target="_blank">
      <Text style="max-width: 54px;font-size:12px;" color="var(--complementary-white-50)" type="p">How we spend it</Text>
    </a>
  </div>
</Card>

<style>
  .about {
    display: flex;
    justify-content: space-between;
  }
</style>