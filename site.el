(require 'package)
(package-initialize)
(require 'ox-publish)
(require 'ox-twbs)

(setq org-publish-project-alist
      '(("my-org-site"
         :recursive t
         :base-directory "/Users/cthulhu/Projects/Org-Site/Development/"
         :publishing-directory "/Users/cthulhu/Projects/Org-Site/Production/Src/"
         :publishing-function org-twbs-publish-to-html
         :with-sub-superscript nil
         :auto-sitemap t
         )
        )
      )

(org-publish-all t)

(message "Build Complete")
