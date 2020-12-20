<?php
require_once "utils.php";


function edit_tmpl() {
    $user_info = new Upage();
    if ($_SERVER['REQUEST_METHOD'] === 'POST') {
        $user_info->set($_POST["key"], $_POST["value"]);
    }
    ?>
        <h1>Edit uPage</h1>
        <form method="post">
            <label>Your Favorite CTF: </label>
            <br>
            <input name="key" value="upage_fav_ctf" hidden>
            <input name="value" type="text" placeholder="WeCTF" value="<?php echo $user_info->upage_fav_ctf ?>">
            <input type="submit" class="button button-primary">
        </form>
        <hr>
        <form method="post">
            <label>Your CTFtime Ranking: </label>
            <br>
            <input name="key" value="upage_ranking" hidden>
            <input name="value" type="number" value="<?php echo $user_info->upage_ranking ?>">
            <input type="submit" class="button button-primary">
        </form>
        <hr>
        <form method="post">
            <label>Your Motto: </label>
            <br>
            <input name="key" value="upage_motto" hidden>
            <textarea name="value" type="text" style="width: 90%"><?php echo $user_info->upage_motto ?></textarea>
            <br>
            <input type="submit" class="button button-primary">
        </form>
        <hr>
        <form method="post">
            <label>Your Favorite Song: </label>
            <br>
            <input name="key" value="upage_fav_song" hidden>
            <input name="value" type="text" placeholder="Lost River"
                   value="<?php echo $user_info->upage_fav_song ?>">
            <input type="submit" class="button button-primary">
        </form>
        <hr>
        <form method="post">
            <label>Your Favorite Person: </label>
            <br>
            <input name="key" value="upage_fav_person" hidden>
            <input name="value" type="text" placeholder="shou" value="<?php echo $user_info->upage_fav_person ?>">
            <input type="submit" class="button button-primary">
        </form>
    <?php
}

function my_tmpl() {
    $user_info = new Upage();
    ?>
    <h1>My uPage</h1>
    <strong>Your Favorite CTF: </strong>
    <?php echo $user_info->upage_fav_ctf ?>
    <hr>
    <strong>Your CTFtime Ranking: </strong>
    <?php echo $user_info->upage_ranking ?>
    <hr>
    <strong>Your Motto: </strong>
    <br>
    <?php echo $user_info->upage_motto ?>
    <hr>
    <strong>Your Favorite Song: </strong>
    <?php echo $user_info->upage_fav_song ?>
    <hr>
    <strong>Your Favorite Person: </strong>
    <?php echo $user_info->upage_fav_person ?>
    <?php
}
