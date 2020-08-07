import { storeSeenReply, seenReplies, hidden } from "../state"
import * as options from "../options";
import lang from "../lang"
import { thumbPath, Post } from "../posts"
import { repliedToMe } from "./tab"
import * as util from "../util"
import { View } from "../base"

// Notify the user that one of their posts has been replied to
export default function notifyAboutReply(post: Post) {
	if (seenReplies.has(post.id) || hidden.has(post.id)) {
		return
	}
	storeSeenReply(post.id, post.op)
	if (!document.hidden) {
		return
	}
	repliedToMe(post)

	if (!options.canNotify()) {
		return
	}

	const opts = options.notificationOpts();
	if (options.canShowImages() && post.image) {
		const { sha1, thumb_type: thumbType, spoiler } = post.image;
		if (spoiler) {
			opts.icon = '/assets/spoil/default.jpg';
		} else {
			opts.icon = thumbPath(sha1, thumbType);
		}
	}
	opts.body = 'On board /${post.board}/:\n${post.body}';
	opts.data = post.id; // Persist target, even if browser tab closed
	const n = new Notification(lang.ui["quoted"], opts)
	n.onclick = function () {
		this.close();
		window.focus();
		location.hash = "#p" + this.data;
		util.scrollToAnchor();
	};
}

// Textual notification at the top of the page
export class OverlayNotification extends View<null> {
	constructor(text: string) {
		super({
			el: util.importTemplate("notification")
				.firstChild as HTMLElement,
		})
		this.on("click", () =>
			this.remove())
		this.el.querySelector("b").textContent = text

		const cont = document.getElementById("modal-overlay");
		let last: HTMLElement;
		for (let i = cont.children.length - 1; i >= 0; i--) {
			const el = cont.children[i];
			if (el.classList.contains("notification")) {
				last = el as HTMLElement;
				break;
			}
		}
		if (last) {
			last.after(this.el);
		} else {
			cont.prepend(this.el);
		}
	}
}
