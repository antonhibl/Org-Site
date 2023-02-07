(require 'ox-publish)

(setq org-publish-project-alist
      (list
       (list "org-site"
             :recursive t
             :base-directory "/Users/cthulhu/Projects/Org-Site/Development/"
             :publishing-directory "/Users/cthulhu/Projects/Org-Site/Production/Src/"
             :publishing-function 'org-twbs-export-to-html
             :with-sub-superscript nil)))

(org-publish-all t)

(message "Build Complete")
